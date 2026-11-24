"use client";
import ComponentSymbol from "./component-symbol";
import { useEffect, useRef, useState } from "react";
import {
  startingSystem,
  traceOrder,
  clients,
  services,
  type Block,
  type Component,
  type Link,
} from "./diagram-model";

const width = 800,
  height = 420,
  tileWidth = 136,
  tileHeight = 72;
const clamp = (n: number, max: number) => Math.min(max, Math.max(0, n));
const label = (block: Block) => block.kind + " " + block.id;

export default function Diagram({ scenario }: { scenario: string }) {
  const [initial] = useState(() => startingSystem(scenario));
  const [blocks, setBlocks] = useState<Block[]>(initial.blocks);
  const [links, setLinks] = useState<Link[]>(initial.links);
  const [source, setSource] = useState<number | null>(null);
  const [result, setResult] = useState("");
  const [attempts, setAttempts] = useState(0);
  const viewport = useRef<HTMLDivElement>(null);
  const [fit, setFit] = useState(true);
  const [available, setAvailable] = useState(width);
  useEffect(() => {
    const element = viewport.current;
    if (!element) return;
    const observer = new ResizeObserver(([entry]) => {
      if (entry) setAvailable(entry.contentRect.width);
    });
    observer.observe(element);
    return () => observer.disconnect();
  }, []);
  const scale = fit ? Math.min(1, available / width) : 1;
  const [motionPaused, setMotionPaused] = useState(false);
  const next = useRef(initial.blocks.length + 1);
  const drag = useRef<{
    id: number;
    x: number;
    y: number;
    startX: number;
    startY: number;
    moved: boolean;
  } | null>(null);

  function add(kind: Component) {
    const id = next.current++;
    setBlocks([
      ...blocks,
      {
        id,
        kind,
        x: 18 + ((id - 1) % 5) * 154,
        y: 24 + (Math.floor((id - 1) / 5) % 5) * 75,
      },
    ]);
    setResult("");
  }
  function select(id: number) {
    if (source === null) return setSource(id);
    if (
      id !== source &&
      !links.some((link) => link.from === source && link.to === id)
    ) {
      setLinks([...links, { from: source, to: id, enabled: true }]);
      setResult("");
    }
    setSource(null);
  }
  function move(id: number, x: number, y: number) {
    setBlocks((current) =>
      current.map((b) =>
        b.id === id
          ? {
              ...b,
              x: clamp(x, width - tileWidth),
              y: clamp(y, height - tileHeight),
            }
          : b,
      ),
    );
  }
  function remove(link: Link) {
    setLinks(links.filter((l) => l !== link));
    setResult("");
  }
  function toggle(target: Link) {
    setLinks((current) =>
      current.map((link) =>
        link === target ? { ...link, enabled: !link.enabled } : link,
      ),
    );
    setResult("");
  }
  const trace = scenario === "sqs-backlog" ? traceOrder(blocks, links) : null;
  const stop = trace?.route.at(-1);
  const packetPath = trace?.route
    .map(
      (node, index) =>
        (index === 0 ? "M" : "L") +
        (node.x + tileWidth / 2) +
        " " +
        (node.y + 26),
    )
    .join(" ");
  const target = trace?.blockedTowards;
  const travelPath =
    stop && target
      ? packetPath +
        " L" +
        (stop.x + target.x + tileWidth) / 2 +
        " " +
        ((stop.y + target.y) / 2 + 26)
      : packetPath;
  const named = (id: number) => blocks.find((b) => b.id === id)!;
  return (
    <article className="diagram">
      <div className="canvas-head">
        <h2>System canvas</h2>
        {stop && (
          <button onClick={() => setMotionPaused(!motionPaused)}>
            {motionPaused ? "Resume animation" : "Pause animation"}
          </button>
        )}
      </div>
      <details className="toolbox">
        <summary>Add components</summary>
        <p>
          Select two nodes in delivery order to connect them. Drag or use arrow
          keys to move a node.
        </p>
        {(
          [
            ["Clients", clients],
            ["Infrastructure", services],
          ] as const
        ).map(([group, kinds]) => (
          <fieldset className="component-group" key={group}>
            <legend>{group}</legend>
            <div className="palette">
              {kinds.map((kind) => (
                <button key={kind} onClick={() => add(kind)}>
                  <ComponentSymbol kind={kind} /> {kind}
                </button>
              ))}
            </div>
          </fieldset>
        ))}
      </details>
      <div
        className="board-scroll"
        ref={viewport}
        role="region"
        aria-label="Scrollable system diagram"
        tabIndex={0}
      >
        <div
          style={{
            width: width * scale,
            height: height * scale,
            marginInline: "auto",
            overflow: "hidden",
          }}
        >
          <div
            className="board"
            style={{
              width,
              height,
              transformOrigin: "top left",
              transform: `scale(${scale})`,
            }}
          >
            <svg width={width} height={height} aria-hidden="true">
              <defs>
                <marker
                  id="direction"
                  viewBox="0 0 10 10"
                  refX="8"
                  refY="5"
                  markerWidth="8"
                  markerHeight="8"
                  orient="auto"
                >
                  <path d="M0 0L10 5L0 10Z" fill="context-stroke" />
                </marker>
              </defs>
              {links.map((link) => {
                const a = named(link.from),
                  b = named(link.to);
                const dx = b.x - a.x,
                  dy = b.y - a.y;
                const scale = Math.min(
                  1,
                  1 / Math.max(Math.abs(dx) / 29, Math.abs(dy) / 26),
                );
                return (
                  <line
                    key={link.from + "-" + link.to}
                    x1={a.x + tileWidth / 2}
                    y1={a.y + 26}
                    x2={b.x + tileWidth / 2 - dx * scale}
                    y2={b.y + 26 - dy * scale}
                    stroke={!link.enabled ? "#af4227" : "#47796a"}
                    strokeWidth="2"
                    strokeDasharray={!link.enabled ? "5 4" : undefined}
                    markerEnd="url(#direction)"
                  />
                );
              })}
              {trace && packetPath && (
                <path
                  className="order-trail"
                  d={packetPath}
                  fill="none"
                  stroke="#3269b6"
                  strokeWidth="4"
                  strokeDasharray="3 9"
                />
              )}
            </svg>
            {trace && (
              <div
                className="flow-note"
                data-delivered={trace.delivered}
                style={{
                  left: stop
                    ? Math.max(85, Math.min(width - 85, stop.x + tileWidth / 2))
                    : width / 2,
                  top: stop
                    ? stop.y < 40
                      ? stop.y + tileHeight + 4
                      : stop.y - 34
                    : 18,
                }}
              >
                <b aria-hidden="true">{trace.delivered ? "✓" : "!"}</b>
                {trace.caption}
              </div>
            )}
            {trace &&
              travelPath &&
              [0, 1, 2].map((index) => (
                <span
                  key={attempts + "-" + trace.delivered + "-" + index}
                  aria-hidden="true"
                  className="order-packet"
                  data-broken={!trace.delivered}
                  data-paused={motionPaused}
                  style={{
                    offsetPath: `path('${travelPath}')`,
                    animationDelay: -index * 1.4 + "s",
                  }}
                >
                  <i />
                  <i />
                </span>
              ))}
            {links.map((link) => {
              const a = named(link.from),
                b = named(link.to);
              return (
                <button
                  key={link.from + "-" + link.to}
                  className="edge-control"
                  data-enabled={link.enabled}
                  style={{
                    left: (a.x + b.x + tileWidth) / 2,
                    top: (a.y + b.y) / 2 + 26,
                  }}
                  aria-label={`${link.enabled ? "Pause" : "Resume"} delivery from ${label(a)} to ${label(b)}`}
                  title={link.enabled ? "Pause delivery" : "Resume delivery"}
                  onClick={() => toggle(link)}
                >
                  {link.enabled ? "Ⅱ" : "▶"}
                </button>
              );
            })}
            {blocks.map((b) => (
              <button
                key={b.id}
                className="tile"
                style={{
                  left: b.x,
                  top: b.y,
                  width: tileWidth,
                  height: tileHeight,
                }}
                aria-pressed={source === b.id}
                onPointerDown={(e) => {
                  drag.current = {
                    id: b.id,
                    x: b.x,
                    y: b.y,
                    startX: e.clientX,
                    startY: e.clientY,
                    moved: false,
                  };
                  e.currentTarget.setPointerCapture(e.pointerId);
                }}
                onPointerCancel={() => {
                  drag.current = null;
                }}
                onPointerMove={(e) => {
                  const d = drag.current;
                  if (!d || d.id !== b.id || !e.buttons) return;
                  const dx = e.clientX - d.startX,
                    dy = e.clientY - d.startY;
                  if (Math.abs(dx) + Math.abs(dy) > 4) d.moved = true;
                  if (d.moved) move(b.id, d.x + dx / scale, d.y + dy / scale);
                }}
                onClick={(e) => {
                  if (!e.detail || !drag.current?.moved) select(b.id);
                  drag.current = null;
                }}
                onKeyDown={(e) => {
                  const offsets: Record<string, [number, number]> = {
                    ArrowLeft: [-10, 0],
                    ArrowRight: [10, 0],
                    ArrowUp: [0, -10],
                    ArrowDown: [0, 10],
                  };
                  const delta = offsets[e.key];
                  if (delta) {
                    e.preventDefault();
                    move(b.id, b.x + delta[0], b.y + delta[1]);
                  }
                }}
              >
                <ComponentSymbol kind={b.kind} />
                <span>{label(b)}</span>
              </button>
            ))}
          </div>
        </div>
      </div>
      {available < width && (
        <button className="zoom" onClick={() => setFit(!fit)}>
          {fit ? "Zoom to edit" : "Fit system to screen"}
        </button>
      )}
      {source !== null && (
        <div className="connection">
          Connecting from {label(named(source))}
          <button onClick={() => setSource(null)}>Cancel</button>
          <button
            onClick={() => {
              setBlocks(blocks.filter((b) => b.id !== source));
              setLinks(
                links.filter((l) => l.from !== source && l.to !== source),
              );
              setSource(null);
              setResult("");
            }}
          >
            Delete component
          </button>
        </div>
      )}
      {links.map((l) => (
        <div className="connection" key={l.from + "-" + l.to}>
          <span>
            {label(named(l.from))} → {label(named(l.to))} ·{" "}
            {l.enabled ? "Enabled" : "Disabled"}
          </span>
          <button
            aria-label={`${l.enabled ? "Disable" : "Enable"} ${label(named(l.from))} to ${label(named(l.to))}`}
            onClick={() => toggle(l)}
          >
            {l.enabled ? "Disable" : "Enable"}
          </button>
          <button onClick={() => remove(l)}>Disconnect</button>
        </div>
      ))}
      {scenario === "sqs-backlog" && (
        <div className="connection">
          <button
            className="check-design"
            onClick={() => {
              setAttempts((n) => n + 1);
              setResult(trace!.message);
            }}
          >
            Send test order
          </button>
          <button
            onClick={() => {
              setAttempts(0);
              setBlocks(initial.blocks);
              setLinks(initial.links);
              next.current = initial.blocks.length + 1;
              setSource(null);
              setResult("Original incident restored.");
            }}
          >
            Restart puzzle
          </button>
        </div>
      )}
      <div
        className="receipt"
        role="status"
        data-success={result.startsWith("Puzzle solved")}
      >
        {result && (
          <>
            <strong>{attempts ? "Test run " + attempts : "Canvas"}</strong>
            <p>{result}</p>
          </>
        )}
      </div>
    </article>
  );
}
