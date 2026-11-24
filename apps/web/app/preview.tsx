"use client";
import { useState } from "react";
import ComponentSymbol from "./practice/[id]/component-symbol";

const stations = [
  {
    kind: "Web browser",
    name: "Orders",
    detail: "Clients continue submitting new orders.",
  },
  {
    kind: "SQS queue",
    name: "Queue",
    detail: "At 10:05 UTC, the oldest order has waited five minutes.",
  },
  {
    kind: "Lambda worker",
    name: "Fulfillment",
    detail: "The worker recorded zero successful orders in five minutes.",
  },
] as const;

export default function Preview() {
  const [active, setActive] = useState(1);
  return (
    <div className="preview">
      <div className="trace-head">
        <strong>Order delivery path</strong>
        <span className="signal">Backlog</span>
      </div>
      <div className="stations">
        {stations.map((s, i) => (
          <button
            key={s.kind}
            aria-pressed={active === i}
            onClick={() => setActive(i)}
          >
            <ComponentSymbol kind={s.kind} />
            <span>{s.name}</span>
          </button>
        ))}
      </div>
      <p aria-live="polite">{stations[active]?.detail}</p>
    </div>
  );
}
