Nezha - We don't trust!
-----------------------

Nezha is an application deployment and secure access platform based on the Zero-Trust security framework for any application, irrespective of scale, complexity, and user size.

Installation
------------

Nezha requires a Kubernetes environment running Kubectl and Helm, the command line tool and the package manager for Kubernetes.

Run your own Nezha server
+++++++++++++++++++++++++++++++
On a k8s cluster install and configure Nezha server using the following instructions.
NOTE: These installation steps assume that you have Helm and kubectl - the package manager and command line tool for k8s - installed and setup already.

    1.  Add Nezha's Helm repository
        ::
            helm repo add nezha-helm https://arorasoham9.github.io/nezha-helm/

    2.  Install Nezha Server
        ::
            helm install Nezha-server nezha-helm/Nezha-server

    3.  Get the loadbalancer external IP for the client to connect to
        ::
            kubectl get svc -n Nezha-helm
        Save the external IP to the service named "Nezha-helm-port-forwarding" for later use, when setting up the Nezha client.
    You may uninstall Nezha Server using the following command::

        helm uninstall Nezha-server

Setup Nezha Client CLI
_____________________
On the machine, install and configure Nezha Client using the following instructions.
    1.  Download Nezha Client::

            curl -O https://link.storjshare.io/s/juzbdy3atfth2i3tdxagot66ujda/nezha/Nezha

    2.  Make it executable
        ::
            chmod +x Nezha*

    3.  Run the below command to learn how to connet to your deployed applications.
        ::
            ./Nezha --help








