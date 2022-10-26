#template example test

#### Test

```bash
$ go test -v
=== RUN   Test_tmpFunc
=== RUN   Test_tmpFunc/tmp_test_

	apiVersion: 1.0.1
	kind: ConfigMap
	metadata:
  		name: One-configmap
	data:
  		myvalue: Lang .C
  === RUN   Test_tmpFunc/tmp_test_#01

	apiVersion: 1.0.1
	kind: ConfigMap
	metadata:
  		name: One-configmap
	data:
  		myvalue: Lang .C
  --- PASS: Test_tmpFunc (0.00s)
    --- PASS: Test_tmpFunc/tmp_test_ (0.00s)
    --- PASS: Test_tmpFunc/tmp_test_#01 (0.00s)
PASS
ok  	tmpl4	0.001s

```
