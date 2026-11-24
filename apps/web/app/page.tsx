import Link from "next/link";
import Masthead from "./masthead";
import Preview from "./preview";
import { listIncidents } from "./catalog";

export const dynamic = "force-dynamic";
export default async function Library() {
  const items = await listIncidents().catch(() => null);
  return (
    <>
      <Masthead />
      <main className="library">
        <div className="intro">
          <h1>Incident library</h1>
          <p>
            Walk into a broken system. Use the evidence, change the
            architecture, and test your repair.
          </p>
        </div>
        {items === null ? (
          <p role="alert">
            The catalog is unavailable. <a href="/">Retry</a>
          </p>
        ) : items.length === 0 ? (
          <p>No exercises available yet.</p>
        ) : (
          items.map((item) => (
            <article className="case" key={item.id}>
              <div className="case-copy">
                <small>{item.platform.toUpperCase()} investigation</small>
                <h2>{item.title}</h2>
                {item.id === "sqs-backlog" && (
                  <p>
                    Orders are arriving. Fulfillment has stopped. Find where
                    delivery breaks.
                  </p>
                )}
                <Link
                  className="primary"
                  href={"/practice/" + encodeURIComponent(item.id)}
                >
                  Open investigation <span aria-hidden="true">↗</span>
                </Link>
              </div>
              {item.id === "sqs-backlog" && <Preview />}
            </article>
          ))
        )}
      </main>
    </>
  );
}
