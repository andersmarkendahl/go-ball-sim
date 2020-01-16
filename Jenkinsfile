pipeline {
    agent none
    environment {
        SLAVE_INFO = 'echo $SHELL, $USER, $PATH'
    }
    stages {
        stage('Smoke Test') {
            parallel {
                stage('Compile') {
                    agent {
                        label "go-compiler"
                    }
                    steps {
                        sh "$SLAVE_INFO"
                        sh "go build -o bounce cmd/bounce/main.go"
                    }
                }
                stage('Test') {
                    agent {
                        label "go-tester"
                    }
                    steps {
                        wrap([$class: 'Xvfb']) {
                            sh "$SLAVE_INFO"
                            sh "go test ./pkg/..."
                        }
                    }
                }
                stage('Syntax') {
                    agent {
                        label "go-syntaxer"
                    }
                    steps {
                        sh "$SLAVE_INFO"
                        sh "golint -set_exit_status ./..."
                        sh "test -z `gofmt -l .`"
                    }
                }
            }
        }
    }
}
