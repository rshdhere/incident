# AWS incident practice

Use the synthetic evidence in this directory for a tabletop interview. The exercise runs on paper or in a text editor.

## Orders stop reaching fulfillment

Open the [scenario](sqs-backlog.json). Give the learner the brief and question first. Reveal each evidence item when they ask to inspect that source. Ask for a diagnosis, a mitigation and a recovery signal before discussing the worked response.

## Author another exercise

Copy the JSON shape into a file named for the new scenario ID. Give each observation a UTC time, source and unit where relevant. Pair SQS metrics with evidence from the Lambda consumer. Use fictional service names and synthetic observations.

Check the service behavior against the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/troubleshooting-event-source-mapping.html).
