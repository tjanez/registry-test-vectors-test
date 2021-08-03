# Example Validation of Oasis Core Entity Descriptor Signatures from Registry's Test Vectors

The example is in `register-entity.json` is taken from [Oasis Core Registry's
Test Vectors][registry-vectors].

[registry-vectors]:
  https://docs.oasis.dev/oasis-core/high-level-components/index/services/registry#test-vectors

## Instructions

```bash
cd $(mktemp -d)
git clone https://github.com/tjanez/registry-test-vectors-test.git
cd registry-test-vectors-test/
go run main.go
```

## Expected results

```
Entity: <Entity id=3wNS/vFr/qqy6oAqgBzesWUMhZB7C8DnCED4T/NKy6M=>
SUCCESS: Entity descriptor signature matches!
```
