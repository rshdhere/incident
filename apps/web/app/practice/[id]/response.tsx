"use client";
import { useState } from "react";

type Answer = { diagnosis: string; mitigation: string; recovery: string };

export default function Response({ prompt, reveal }: { prompt: string; reveal: () => Promise<Answer> }) {
  const [answer, setAnswer] = useState<Answer | null>(null);
  const [pending, setPending] = useState(false);
  const [failed, setFailed] = useState(false);
  async function compare() {
    setPending(true);
    setFailed(false);
    try { setAnswer(await reveal()); }
    catch { setFailed(true); }
    finally { setPending(false); }
  }
  return <article><h2>Your investigation</h2>
    <label htmlFor="notes">{prompt}</label>
    <textarea id="notes" rows={4} placeholder="Write your reasoning…" />
    {answer ? <section aria-label="Worked response"><h2>Worked response</h2>
      {Object.entries(answer).map(([key,value]) => <p key={key}><b className="term">{key}</b>{value}</p>)}
    </section> : <button disabled={pending} onClick={compare}>{pending ? "Loading…" : "Reveal worked response"}</button>}
    {failed && <p role="alert">The response could not be loaded. Try again.</p>}
  </article>;
}
