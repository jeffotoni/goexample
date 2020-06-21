gcloud functions deploy Pub \
  --runtime go113 \
  --trigger-http \
  --allow-unauthenticated

Deploying function (may take a while - up to 2 minutes)...done.                
availableMemoryMb: 256
entryPoint: Gopher
httpsTrigger:
  url: https://region-my-project.cloudfunctions.net/Pub
...