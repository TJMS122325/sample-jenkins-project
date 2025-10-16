pipeline {
    agent any

    environment {
        // Change these values
        APP_NAME = "my-app"
        DOCKERHUB_USER = "TJMS122325"
        IMAGE_NAME = "${DOCKERHUB_USER}/${APP_NAME}:latest"
        GIT_REPO = "https://github.com/TJMS122325/sample-jenkins-project.git"
    }

    stages {
        stage('Checkout') {
            steps {
                echo "Checking out code from GitHub..."
                git branch: 'main', url: "${GIT_REPO}"
            }
        }

        stage('Build') {
            steps {
                echo "Building application..."
                sh '''
                # Example for a Go app — adjust for your stack
                go mod tidy
                go build -o ${APP_NAME}
                '''
            }
        }

        stage('Test') {
            steps {
                echo "Running tests..."
                sh '''
                go test ./... -v
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Building Docker image..."
                sh '''
                docker build -t ${IMAGE_NAME} .
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
                echo "Pushing Docker image to Docker Hub..."
                withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                    echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                    docker push ${IMAGE_NAME}
                    docker logout
                    '''
                }
            }
        }
    }

    post {
        success {
            echo "✅ CI/CD pipeline completed successfully!"
        }
        failure {
            echo "❌ Pipeline failed. Check the logs for details."
        }
    }
}
