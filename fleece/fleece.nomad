job "fleece.liquidity" {
  region = "global"

  datacenters = ["dc1"]

  type = "service"

  group "x.liquidity.types" {
    count = 1

    #volume "liquidity" {
    #    type = "host"
    #    source = "liquidity"
    #    read_only = false
    #}

    task "FuzzMsgSwapWithinBatch_raw" {
        #volume_mount {
        #    volume = "liquidity"
        #    destination = "/tmp/fuzzing"
        #    read_only = false
        #}

        driver = "docker"
        config {
            # TODO: setup authenticated registry.
            # TODO: push image to authenticated registry.
            image = "bryanchriswhite/go-fuzz:latest"
            #command = "bash"
            #auth {
            #   server_address = "060199087156.dkr.ecr.eu-central-1.amazonaws.com"
            #   # NB: user only has read permission on /liquidity-fuzzing repo.
            #   username = "AKIAQ4BBFFQ2MHH7GO6M"
            #   password = "KcQFjzulC9mPxvLGhgbBRsyLKdk9tvIGJqis/jIB"
            #}
            entrypoint = ["/go-fuzz.sh"]
            #args = ["-c", "ls && sleep infinity"]
            args = [
                "./x/liquidity/types", "FuzzMsgSwapWithinBatch_raw",
                "--", "-workdir", "./alloc/liquidity/fleece/workdirs/FuzzMsgSwapWithinBatch_raw",
            ]
            volumes = ["/home/bwhite/Projects/liquidity:/tmp/fuzzing"]
            #mount {
            #    type = "volume"
            #    target = "/tmp/fuzzing"
            #    source = "liquidity"
            #    readonly = false
            #}
        }

        resources {
            # TODO: may make more sense to define in terms of cpu
            # after making separate worker task
            cpu = 100
            #cores = 1 # 5
            memory = 256
            #memory_max = 2000
        }
    }
    task "FuzzMsgSwapWithinBatch_structured" {
        driver = "docker"
        config {
            image = "bryanchriswhite/go-fuzz:latest"
            entrypoint = ["/go-fuzz.sh"]
            args = [
                "./x/liquidity/types", "FuzzMsgSwapWithinBatch_structured",
                "--", "-workdir", "./alloc/liquidity/fleece/workdirs/FuzzMsgSwapWithinBatch_structured",
            ]
            volumes = ["/home/bwhite/Projects/liquidity:/tmp/fuzzing"]
        }
        resources {
            # TODO: may make more sense to define in terms of cpu
            # after making separate worker task
            cpu = 100
            #cores = 1 # 5
            memory = 256
            #memory_max = 2000
        }
    }
    task "FuzzMsgWithdrawWithinBatch_raw" {
        driver = "docker"
        config {
            image = "bryanchriswhite/go-fuzz:latest"
            entrypoint = ["/go-fuzz.sh"]
            args = [
                "./x/liquidity/types", "FuzzMsgWithdrawWithinBatch_raw",
                "--", "-workdir", "./alloc/liquidity/fleece/workdirs/FuzzMsgWithdrawWithinBatch_raw",
            ]
            volumes = ["/home/bwhite/Projects/liquidity:/tmp/fuzzing"]
        }
        resources {
            # TODO: may make more sense to define in terms of cpu
            # after making separate worker task
            cpu = 100
            #cores = 1 # 5
            memory = 256
            #memory_max = 2000
        }
    }
    task "FuzzMsgWithdrawWithinBatch_structured" {
        driver = "docker"
        config {
            image = "bryanchriswhite/go-fuzz:latest"
            entrypoint = ["/go-fuzz.sh"]
            args = [
                "./x/liquidity/types", "FuzzMsgWithdrawWithinBatch_structured",
                "--", "-workdir", "./alloc/liquidity/fleece/workdirs/FuzzMsgWithdrawWithinBatch_raw",
            ]
            volumes = ["/home/bwhite/Projects/liquidity:/tmp/fuzzing"]
        }
        resources {
            # TODO: may make more sense to define in terms of cpu
            # after making separate worker task
            cpu = 100
            #cores = 1 # 5
            memory = 256
            #memory_max = 2000
        }
    }
  }
}
