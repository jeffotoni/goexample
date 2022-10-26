#template example test

#### Test

```bash
$ go test -v
=== RUN   Test_tmpFunc
=== RUN   Test_tmpFunc/tmp_test_01

	apiVersion: 1.0.0.0
	kind: ConfigMap
	metadata:
  		name: One-configmap
	data:
  		myvalue: Lang C
  === RUN   Test_tmpFunc/tmp_test_02

	apiVersion: 2.0.0.0
	kind: ConfigMap
	metadata:
  		name: Two-configmap
	data:
  		myvalue: Lang Go
  --- PASS: Test_tmpFunc (0.00s)
    --- PASS: Test_tmpFunc/tmp_test_01 (0.00s)
    --- PASS: Test_tmpFunc/tmp_test_02 (0.00s)
PASS
ok  	tmpl4	0.001s

```
