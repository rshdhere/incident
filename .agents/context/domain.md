# Scenario vocabulary

Use these terms when discussing the planned learner experience. They describe content and behavior. Storage schemas and UI layouts remain open.

| Term | Meaning in an exercise |
| --- | --- |
| Incident | A simulated service failure with observable impact and a timeline. |
| Scenario | The authored interview exercise that presents the incident and prompts the learner. |
| Evidence | Something the learner can inspect, such as a metric, deployment event or log excerpt. |
| Hypothesis | A possible explanation the learner can test against the evidence. |
| Canvas | The learner's guided workspace for inspecting evidence and explaining the investigation. |
| Mitigation | An action intended to reduce impact while the cause is investigated or repaired. |
| Recovery signal | An observation that supports the claim that service has recovered. |

For an AWS exercise involving a surge in queue depth, the queue graph is evidence. A stalled consumer is a hypothesis. Healthy processing rates and falling message age can help verify a mitigation. Write prompts that reward the reasoning connecting those steps.
