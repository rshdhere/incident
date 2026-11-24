# Azure incident practice

Use the synthetic evidence in this directory for a tabletop interview. The exercise runs on paper or in a text editor.

## Invoices collect in a dead-letter queue

Open the [scenario](service-bus-deliveries.json). Give the learner the brief and question first. Reveal each evidence item when they ask to inspect that source. Ask for a diagnosis, a mitigation and a recovery signal before discussing the worked response.

## Author another exercise

Copy the JSON shape into a file named for the new scenario ID. Give each observation a UTC time, source and unit where relevant. Connect broker-side signals to consumer-side observations. Use fictional service names and synthetic observations.

Check the service behavior against the [Azure documentation](https://learn.microsoft.com/en-us/azure/service-bus-messaging/service-bus-dead-letter-queues).
