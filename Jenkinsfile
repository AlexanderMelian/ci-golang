pipeline {
    agent any

    options {
        timestamps()
        disableConcurrentBuilds()
    }

    stages {
        stage('Info') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                sh 'go version'
            }
        }

        stage('Format') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                sh 'test -z "$(gofmt -l .)"'
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                sh 'go build ./...'
            }
        }

        stage('Test') {
            agent {
                docker {
                    image 'golang:1.26'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    go install gotest.tools/gotestsum@latest
                    /go/bin/gotestsum --junitfile junit.xml -- -v ./...
                '''
            }
            post {
                always {
                    junit testResults: 'junit.xml', allowEmptyResults: true
                }
            }
        }
    }

    post {
        success {
            echo 'Build y tests OK'
        }
        failure {
            echo 'Falló el pipeline'
        }
    }
}