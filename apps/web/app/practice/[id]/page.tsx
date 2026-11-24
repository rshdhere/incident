import Masthead from "../../masthead";
import { getExercise, getDebrief } from "../../catalog";
import Response from "./response";
import Diagram from "./diagram";

export default async function Practice({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const exercise = await getExercise(id).catch(() => null);
  async function reveal() {
    "use server";
    return getDebrief(id);
  }
  return (
    <>
      <Masthead practice />
      <main className="practice">
        {exercise ? (
          <>
            <div className="brief">
              <small>{exercise.platform.toUpperCase()} investigation</small>
              <h1>{exercise.title}</h1>
              <p>{exercise.brief}</p>
            </div>
            <div className="workbench">
              <Diagram key={exercise.id} scenario={exercise.id} />
              <aside className="dossier">
                <h2>Evidence timeline</h2>
                {exercise.evidence.map((item) => (
                  <details key={item.time + item.source}>
                    <summary>
                      <time>{item.time}</time> {item.source}
                    </summary>
                    <p>{item.observation}</p>
                  </details>
                ))}

                <Response prompt={exercise.question} reveal={reveal} />
              </aside>
            </div>
          </>
        ) : (
          <section>
            <h1>Exercise unavailable</h1>
            <a href="">Try again</a>
          </section>
        )}
      </main>
    </>
  );
}
