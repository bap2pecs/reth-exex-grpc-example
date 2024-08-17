use futures::Future;
use reth_exex::{ExExContext, ExExEvent, ExExNotification};
use reth_node_api::FullNodeComponents;
use reth_node_ethereum::EthereumNode;
use reth_tracing::tracing::info;

use tonic::transport::Channel;
use tonic::Request;

// Import the generated client code
mod exex {
    tonic::include_proto!("exex");
}

use exex::ex_ex_handler_client::ExExHandlerClient;
use exex::ChainCommittedRequest;


/// The initialization logic of the ExEx is just an async function.
///
/// During initialization you can wait for resources you need to be up for the ExEx to function,
/// like a database connection.
async fn exex_init<Node: FullNodeComponents>(
    ctx: ExExContext<Node>,
) -> eyre::Result<impl Future<Output = eyre::Result<()>>> {
    Ok(exex(ctx))
}

/// An ExEx is just a future, which means you can implement all of it in an async function!
///
/// This ExEx just prints out whenever either a new chain of blocks being added, or a chain of
/// blocks being re-orged. After processing the chain, emits an [ExExEvent::FinishedHeight] event.
async fn exex<Node: FullNodeComponents>(mut ctx: ExExContext<Node>) -> eyre::Result<()> {
    // Create a channel to the server
    // TODO: make this configurable
    let channel = Channel::from_static("http://host.docker.internal:50051").connect().await?;

    // Create a client
    let mut client = ExExHandlerClient::new(channel);

    while let Some(notification) = ctx.notifications.recv().await {
        match &notification {
            ExExNotification::ChainCommitted { new } => {
                 // Create a request
                let request = Request::new(ChainCommittedRequest {
                    start_block: new.first().number,
                    end_block: new.tip().number,
                });

                // Call the gRPC method
                let response = client.handle_chain_committed(request).await?;
                info!(grpc_response = ?response, "Received commit");

                let inner = response.into_inner();

                // call ExExEvent::FinishedHeight if send_finished_height_event is true in the response
                if inner.send_finished_height_event {
                    let finished_height = inner.finished_height;
                    ctx.events.send(ExExEvent::FinishedHeight(finished_height))?;
                }
            }
            ExExNotification::ChainReorged { old, new } => {
                // TODO: reorg handling using the gRPC client
                info!(from_chain = ?old.range(), to_chain = ?new.range(), "Received reorg");
            }
            ExExNotification::ChainReverted { old } => {
                // TODO: revert handling using the gRPC client
                info!(reverted_chain = ?old.range(), "Received revert");
            }
        };

        // if let Some(committed_chain) = notification.committed_chain() {
        //     ctx.events.send(ExExEvent::FinishedHeight(committed_chain.tip().number))?;
        // }
    }

    Ok(())
}

fn main() -> eyre::Result<()> {
    reth::cli::Cli::parse_args().run(|builder, _| async move {
        let handle = builder
            .node(EthereumNode::default())
            .install_exex("Minimal", exex_init)
            .launch()
            .await?;

        handle.wait_for_node_exit().await
    })
}