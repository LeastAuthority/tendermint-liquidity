job "fleece.liquidity" {
  region = "global"

  datacenters = ["dc1"]

  type = "service"

  group "x.liquidity.types" {
    count = 1

    task "FuzzMsgSwapWithinBatch_raw" {
        driver = "docker"
        config {
            image = "bryanchriswhite/go-fuzz:latest"
            entrypoint = ["/go-fuzz.sh"]
            args = [
                "./x/liquidity/types", "FuzzMsgSwapWithinBatch_raw",
                "--", "-workdir", "./alloc/liquidity/fleece/workdirs/FuzzMsgSwapWithinBatch_raw",
            ]
            volumes = ["/home/bwhite/Projects/liquidity:/tmp/fuzzing"]
            cpu_hard_limit = true
        }
        resources {
            cpu = 2700
            memory = 256
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
            cpu_hard_limit = true
        }
        resources {
            cpu = 2700
            memory = 256
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
            cpu_hard_limit = true
        }
        resources {
            cpu = 2700
            memory = 256
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
            cpu_hard_limit = true
        }
        resources {
            cpu = 2700
            memory = 256
        }
    }
  }
}
