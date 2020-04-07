apiVersion: {{.Version}}
kind: {{.Kind}}
metadata:
  name: {{.Release.Name}}-configmap
data:
  myvalue: {{.Title}}