package yunpian

var TestConfig = DefaultDevConfig().WithAPIKey("")
var TestClient = NewClient(TestConfig)
