node('windows') {
    def projectRepo = 'tranhuucuong91'
    def appName = 'fibo'
    def imageTag = "${projectRepo}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"

    checkout scm

    stage 'Build image'
    sh("docker build -t ${imageTag} .")

    stage 'Run Go tests'
    sh("docker run ${imageTag} go test")

    // stage 'Push image to registry'
    // sh("gcloud docker push ${imageTag}")

    stage "Deploy Application"
    switch (env.BRANCH_NAME) {
    // Roll out to staging
    case "staging":
        // Change deployed image in staging
        sh("echo Deploy in staging environment")
        break

    // Roll out to production
    case "master":
        // Change deployed image in production
        sh("echo Deploy in production environment")
        break

    // Roll out a dev environment
    default:
        // Create namespace if it doesn't exist
        sh("echo Create dev environment")
        sh("echo Deploy in dev environment")
    }
}
