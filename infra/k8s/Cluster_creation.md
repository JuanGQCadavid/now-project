Master:

* Install `K3s -> curl -sfL https://get.k3s.io | sh -`
* Get token -> `sudo cat /var/lib/rancher/k3s/server/node-token`

Workers:
* `curl -sfL https://get.k3s.io | K3S_URL=https://IP_ADDRESS_MASTER:6443 K3S_TOKEN=TOKEN sh -`

Connect to master:
* On master machine - `sudo cat /etc/rancher/k3s/k3s.yaml`
* On laptop: `vim .kube/config` - Paste here an replace 127.0.0.1 for the master ip address
