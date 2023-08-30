//def userId = input message: 'Approve merge to master?', submitter: 'developer', submitterParameter: 'USER'
def userId = input message: 'Approve merge to master?', submitterParameter: 'USER'
println "Accepted by ${userId}"

def user = User.getById(userId, false)

println "User Name: ${user.getDisplayName()}"

def authStrategy = Hudson.instance.getAuthorizationStrategy()
def permissions = authStrategy.roleMaps.inject([:]){map, it -> map + it.value.grantedRoles}
def roles = permissions.findAll{ it.value.contains(userId) }.collect{it.key.name}
println "Roles: ${roles}"

pipeline {
    agent { docker { image 'golang:1.21.0-alpine3.18' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
