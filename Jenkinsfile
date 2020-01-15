pipeline {
    agent none
    stages {
        stage('Smoke Test') {
            parallel {
                stage('Compile') {
                    agent {
                        label "go-compiler"
                    }
                    steps {
                        sh "echo $USER,$PATH"
                        sh "go build -o gravity cmd/gravity/main.go"
                    }
                }
                stage('Test') {
                    agent {
                        label "go-tester"
                    }
                    steps {
                        sh "echo $USER,$PATH"
                        sh "go test ./pkg/..."
                    }
                }
                stage('Syntax') {
                    agent {
                        label "go-syntaxer"
                    }
                    steps {
                        sh "echo $USER,$PATH"
                        sh "golint -set_exit_status ./..."
                        sh "test -z `gofmt -l .`"
                    }
                }
            }
        }
    }
}
