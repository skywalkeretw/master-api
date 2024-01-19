#!/bin/bash
if [ $# -eq 0 ]; then
    echo "Usage: $0 <clustername>"
    exit 1
fi

KIND_CLUSTER_NAME=$1

openssl req \
-newkey rsa:4096 -nodes -sha256 -keyout /tmp/kind-registry.key \
-addext "subjectAltName = DNS:registry.svc.cluster.local" \
-x509 -days 365 -out /tmp/kind-registry.crt -subj "/C=de/ST=bw/L=stuttgart/O=luke/OU=msc/CN=registry.svc.cluster.local"

kubectl create secret tls registry-tls --cert=/tmp/kind-registry.crt --key=/tmp/kind-registry.key
kubectl apply -f deployment/registry.yml
kubectl rollout status deployment registry --timeout=1m

# Enable the KinD cluster to pull from the container registry
echo "[INFO] Enabling the KinD cluster to pull from the container registry"
REGISTRY_IP="$(kubectl get service registry -o json | jq -r '.spec.clusterIP')"
echo "Registry IP: ${REGISTRY_IP}"

for NODE in ${KIND_CLUSTER_NAME}-control-plane; do
    docker exec "${NODE}" bash -c "echo '${REGISTRY_IP} registry.svc.cluster.local' >> /etc/hosts"
    docker cp /tmp/kind-registry.crt "${NODE}":/usr/local/share/ca-certificates/kind-registry.crt
    docker exec "${NODE}" update-ca-certificates
    docker exec "${NODE}" systemctl restart containerd
done
    # 4  vim etc/containerd/config.toml
    # 5  systemctl restart containerd