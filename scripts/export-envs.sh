echo export K8S_CA_CERT="$HOME/.minikube/ca.crt"
echo export K8S_CERT="$HOME/.minikube/profiles/minikube/client.crt"
echo export K8S_KEY_CERT="$HOME/.minikube/profiles/minikube/client.key"
echo export K8S_ADDRESS=$(kubectl config view -o jsonpath='{.clusters[0].cluster.server}')
echo export K8S_JWT_TOKEN=$(kubectl -n kube-system create token default)