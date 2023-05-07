
Nezha - We don't trust!
-----------------------

Nezha is an application deployment and secure access platform based on the Zero-Trust security framework for any application, irrespective of scale, complexity, and user size.

Access and Usage
----------------

We provide you three ways to ways to access your deployed applications or machines via Nezha.

Using the Nezha Client CLI
+++++++++++++++++++++++++++++++
NOTE: The following instructions assume that you have a deployment running somewhere in a Kubernetes environment using the Nezha Helm chart provided in our repository
and have available a loadbalancer HOST:IP or Domain available to connect to and have downloaded the Nezha client CLI. If not, follow the instructions in our readME to get started.

    ./Nezha login <LoadBalancer HOST:IP | Domain>

This is a one time command and you would not be required to run it again provided you do not delete any config files or want to login as another user.
Answer the prompts as they appear and you should be set up to connect to your applications. Run the following command to list your currently running applications.

    ./Nezha list

Run the following command to connect to a particular application

    ./Nezha <application identification ID>

Using the Web Terminal
++++++++++++++++++++++

Using the loadbalancer HOST:IP or domain generatred during server setup go to the following route and follow the instructions as they appear::

    https://<LoadBalancer HOST:IP |Domain>:<PORT>/portal/

Using the following API endpoints
+++++++++++++++++++++++++++++++++

We provide a number of and everincreasing API endpoints to allow you to incorporate Nezha in your own proprietary software.
NOTE: These endpoints will return a non HTTP Status Code 503-Service Unavailable response if any authentication TOKENS or CERTS are missing
