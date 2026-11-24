# Kubernetes incident practice

Use the synthetic evidence in this directory for a tabletop interview. The exercise runs on paper or in a text editor.

## A worker keeps restarting after a rollout

Open the [scenario](worker-restarts.json). Give the learner the brief and question first. Reveal each evidence item when they ask to inspect that source. Ask for a diagnosis, a mitigation and a recovery signal before discussing the worked response.

## Author another exercise

Copy the JSON shape into a file named for the new scenario ID. Give each observation a UTC time, source and unit where relevant. Include container state, pod events or a revision diff that distinguishes the symptom from the cause. Use fictional service names and synthetic observations.

Check the service behavior against the [Kubernetes documentation](https://kubernetes.io/docs/tasks/configure-pod-container/assign-memory-resource/).
