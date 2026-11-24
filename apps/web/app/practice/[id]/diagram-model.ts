export const clients = ["Web browser", "Mobile app", "API client"] as const;
export const services = [
  "Load balancer",
  "CDN",
  "Database",
  "Cache",
  "SQS queue",
  "Lambda worker",
] as const;
export type Component = (typeof clients)[number] | (typeof services)[number];
export type Block = { id: number; kind: Component; x: number; y: number };
export type Link = { from: number; to: number; enabled: boolean };

const producer = 1,
  orders = 2,
  fulfillment = 3;
export function startingSystem(scenario: string): {
  blocks: Block[];
  links: Link[];
} {
  if (scenario !== "sqs-backlog") return { blocks: [], links: [] };
  return {
    blocks: [
      { id: producer, kind: "Web browser", x: 64, y: 160 },
      { id: orders, kind: "SQS queue", x: 332, y: 160 },
      { id: fulfillment, kind: "Lambda worker", x: 600, y: 160 },
    ],
    links: [
      { from: producer, to: orders, enabled: true },
      { from: orders, to: fulfillment, enabled: false },
    ],
  };
}

export function traceOrder(blocks: Block[], links: Link[]) {
  const flows = (from: number, to: number) =>
    links.some((link) => link.from === from && link.to === to && link.enabled);
  const senders = blocks.filter((block) =>
    clients.some((kind) => kind === block.kind),
  );
  const sender = senders.find((block) => flows(block.id, orders));
  if (!sender) {
    const stalled =
      senders.find((block) =>
        links.some((link) => link.from === block.id && link.to === orders),
      ) ?? senders[0];
    return {
      route: stalled ? [stalled] : [],
      blockedTowards: links.some(
        (link) => link.from === stalled?.id && link.to === orders,
      )
        ? blocks.find((block) => block.id === orders)
        : undefined,
      delivered: false,
      caption: senders.length ? "Stopped before SQS" : "Missing order client",
      message:
        "Test failed. The order did not reach the original orders queue. Inspect the producer connection.",
    };
  }
  const queue = blocks.find((block) => block.id === orders)!;
  const worker = blocks.find((block) => block.id === fulfillment);
  if (!worker || !flows(orders, fulfillment))
    return {
      route: [sender, queue],
      blockedTowards: links.some(
        (link) => link.from === orders && link.to === fulfillment,
      )
        ? worker
        : undefined,
      delivered: false,
      caption: "Orders stuck in queue",
      message:
        "Test failed. The order reached SQS and is still waiting. Inspect how the fulfillment worker receives messages.",
    };
  return {
    route: [sender, queue, worker],
    delivered: true,
    caption: "Delivery path open",
    message:
      "Puzzle solved. The test order reached the orders queue and the fulfillment worker processed it.",
  };
}
