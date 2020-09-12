properties([
  buildDiscarder(
    logRotator(
      artifactDaysToKeepStr: '',
      artifactNumToKeepStr: '',
      daysToKeepStr: '',
      numToKeepStr: '3'
    )
  )
])

pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
      label 'docker'
    }
  }

  environment {
    GIT_SHORT_REV = sh(
      script: 'git rev-parse --short "${GIT_COMMIT}"',
      returnStdout: true
    ).trim()
  }

  stages {
    stage('Setup') {
      steps {
        sh '''
        go mod download
        '''
      }
    }

    stage('Build') {
      steps {
        sh '''
        PACK=1 ./scripts/build.sh
        '''
      }
    }

    stage('Archive') {
      steps {
        archiveArtifacts 'bin/*.tar.gz, bin/*.zip'
      }
    }
  }

  post {
    always {
      cleanWs cleanWhenFailure: false
    }
  }
}