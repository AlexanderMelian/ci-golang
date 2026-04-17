pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Version') {
            steps {
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -v'
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
        always {
            echo 'Pipeline finalizado'
        }
    }
}