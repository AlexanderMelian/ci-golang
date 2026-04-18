pipeline {
    agent {
        docker {
            image 'golang:1.26'
            reuseNode true
        }
    }

    options {
        timestamps()
        disableConcurrentBuilds()
    }

    stages {
        stage('Info') {
            steps {
                sh 'go version'
            }
        }

        stage('Format') {
            steps {
                sh 'test -z "$(gofmt -l .)"'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    go install gotest.tools/gotestsum@latest
                    /go/bin/gotestsum --junitfile junit.xml -- -v ./...
                '''
            }
        }
    }

    post {
        always {
            junit testResults: 'junit.xml', allowEmptyResults: true
        }
        success {
            echo 'Build y tests OK'
        }
        failure {
            echo 'Falló el pipeline'
        }
    }
}