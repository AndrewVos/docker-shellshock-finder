# Usage

Install nsenter:

    docker run --rm jpetazzo/nsenter cat /nsenter > /tmp/nsenter && chmod +x /tmp/nsenter

Run docker-shellshock-finder:

    go get github.com/AndrewVos/docker-shellshock-finder
    sudo docker-shellshock-finder
