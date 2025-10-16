pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                echo '📦 Checking out source code...'
                checkout scm
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.22'     // ✅ Use official Go Docker image
                    args '-v $PWD:/app -w /app/go-app' // optional: mount workspace to /app/go-app
                }
            }
            steps {
                echo '🔧 Building Go application...'
                sh 'go version'
                sh 'go mod tidy'
                sh 'go build -o app'
            }
        }

        stage('Test') {
            agent {
                docker { image 'golang:1.22' }
            }
            steps {
                echo '🧪 Running tests...'
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                echo '🐳 Building Docker image...'
                sh 'docker build -t yourdockerhubusername/sample-jenkins-project .'
            }
        }

        stage('Push to Docker Hub') {
            environment {
                DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials-id')
            }
            steps {
                echo '📤 Pushing image to Docker Hub...'
                sh """
                    echo "$DOCKERHUB_CREDENTIALS_PSW" | docker login -u "$DOCKERHUB_CREDENTIALS_USR" --password-stdin
                    docker push yourdockerhubusername/sample-jenkins-project
                """
            }
        }
    }

    post {
        success {
            echo '✅ Pipeline completed successfully!'
        }
        failure {
            echo '❌ Pipeline failed. Check the logs for details.'
        }
    }
}
