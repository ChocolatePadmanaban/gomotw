---
layout: default
---

# String Template Golang 

- Fill template file with required values 

```
package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {

	testTempalate := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .ServiceName }}
  namespace: {{ .Namespace }}
spec:
  replicas: {{ .InstanceCount }}
  selector:
    matchLabels:
      app: {{ .ServiceName }}
      version: v1
  template:
    metadata:
      labels:
        app: {{ .ServiceName }}
        version: v1
    spec:
      containers:
      - image: {{ .ImageName }}
        args: {{ .Fortioargs }}
        imagePullPolicy: IfNotPresent
        name: {{ .ServiceName }}
        ports:
        - containerPort: 80
        resources:
          limits:
            cpu: {{ .Cpu }}
            memory: {{ .Memory }}
          requests:
            cpu: {{ .Cpu }}
            memory: {{ .Memory }}
---
apiVersion: v1
kind: Service
metadata:
  name: {{ .ServiceName }}
  namespace: {{ .Namespace }}
  labels:
    app: {{ .ServiceName }}
    service: {{ .ServiceName }}
spec:
  ports:
  - name: http-test
    port: {{ .ServicePort }}
    targetPort: 80
  selector:
    app: {{ .ServiceName }}
`

	tmpl, err := template.New("NamespaceTemplate").Parse(testTempalate)
	if err != nil {
		fmt.Println(err, "should not have occured")
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, map[string]interface{}{"Namespace": "test-namespace",
		"ServiceName":  "test-service",
		"InstantCount": 2,
		"ServicePort":  25200,
		"ImageName":    "test-image",
		"Fortioargs":   "['server','-http-port','0.0.0.0:80']",
		"Cpu":          "20m",
		"Memory":       "32Mi"})

	if err != nil {
		fmt.Println(err, "should not have occured")
	}
	tmplString := buffer.String()
	fmt.Println(tmplString)

}

```

- Output should have targetted  filled with respective values 

```
 % go run string_templates.go
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-service
  namespace: test-namespace
spec:
  replicas: <no value>
  selector:
    matchLabels:
      app: test-service
      version: v1
  template:
    metadata:
      labels:
        app: test-service
        version: v1
    spec:
      containers:
      - image: test-image
        args: ['server','-http-port','0.0.0.0:80']
        imagePullPolicy: IfNotPresent
        name: test-service
        ports:
        - containerPort: 80
        resources:
          limits:
            cpu: 20m
            memory: 32Mi
          requests:
            cpu: 20m
            memory: 32Mi
---
apiVersion: v1
kind: Service
metadata:
  name: test-service
  namespace: test-namespace
  labels:
    app: test-service
    service: test-service
spec:
  ports:
  - name: http-test
    port: 25200
    targetPort: 80
  selector:
    app: test-service

```