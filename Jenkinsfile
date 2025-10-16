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
