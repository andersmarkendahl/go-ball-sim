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
                        sh "./scripts/bsg.sh build --clean"
                        sh "./scripts/bsg.sh build"
                    }
                }
                stage('Test') {
                    agent {
                        label "go-tester"
                    }
                    steps {
                        wrap([$class: 'Xvfb']) {
                            sh "$SLAVE_INFO"
                            sh "./scripts/bsg.sh test"
                        }
                    }
                }
                stage('Syntax') {
                    agent {
                        label "go-syntaxer"
                    }
                    steps {
                        sh "$SLAVE_INFO"
                        sh "./scripts/bsg.sh format"
                    }
                }
            }
        }
    }
}
