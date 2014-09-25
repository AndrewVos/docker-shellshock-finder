# Usage

Install nsenter:

    docker run --rm jpetazzo/nsenter cat /nsenter > /tmp/nsenter && chmod +x /tmp/nsenter

Run docker-shellshock-finder:

    go get github.com/AndrewVos/docker-shellshock-finder
    sudo docker-shellshock-finder

If you don't have golang installed on your machines then you might want to just `go build` and `scp` the binary to your machines.
