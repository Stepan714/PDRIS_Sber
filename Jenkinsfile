pipeline {
    agent any

    tools {
        maven 'Maven 3'
    }

    environment {
        SONAR_URL = 'http://sonarqube:9000'
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'lab4_v2', url: 'https://github.com/Stepan714/PDRIS_Sber.git'
                sh 'ls -la'
            }
        }

        stage('Build') {
            steps {
                sh 'mvn clean package'
            }
        }
        stage('SonarQube Analysis') {
            steps {
                withSonarQubeEnv('SonarQube') {
                    sh 'mvn sonar:sonar'
                }
            }
        }

        stage('Allure Report') {
            steps {
                allure([
                    results: [[path: '**/target/allure-results']]
                ])
            }
        }

        stage('Deploy') {
            steps {
                script {
                    if (env.DEPLOY_METHOD == 'ansible') {
                        ansiblePlaybook(
                            playbook: 'ansible/deploy.yml',
                            inventory: 'ansible/inventory'
                        )
                    } else {
                        sh 'docker-compose up -d'
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Build and deployment successful!'
        }
        failure {
            echo 'Build or deployment failed.'
        }
    }
}
