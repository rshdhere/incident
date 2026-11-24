import sqs from "../../../.aws/sqs-backlog.json";

type Summary = { id: string; title: string; platform: string };
type Exercise = Summary & {
  brief: string;
  question: string;
  evidence: { time: string; source: string; observation: string }[];
};
type Debrief = { diagnosis: string; mitigation: string; recovery: string };

async function read<T>(path: string): Promise<T> {
  const origin = process.env.INCIDENT_API_ORIGIN || "http://127.0.0.1:8080";
  const response = await fetch(new URL("/api/v1/scenarios" + path, origin), {
    cache: "no-store",
    signal: AbortSignal.timeout(5000),
  });
  if (!response.ok) throw new Error("Could not load the incident catalog");
  return response.json();
}

// Preview and production can serve the authored exercise without a separate API host.
const bundled = Boolean(process.env.VERCEL) && !process.env.INCIDENT_API_ORIGIN;
const summary = { id: sqs.id, title: sqs.title, platform: "AWS" };
export async function listIncidents() {
  return bundled ? [summary] : read<Summary[]>("");
}
export async function getExercise(id: string) {
  if (!bundled) return read<Exercise>("/" + encodeURIComponent(id));
  if (id !== sqs.id) throw new Error("Unknown incident");
  return { ...summary, brief: sqs.brief, question: sqs.question, evidence: sqs.evidence };
}
export async function getDebrief(id: string) {
  if (!bundled) return read<Debrief>("/" + encodeURIComponent(id) + "/debrief");
  if (id !== sqs.id) throw new Error("Unknown incident");
  return sqs.workedResponse;
}
