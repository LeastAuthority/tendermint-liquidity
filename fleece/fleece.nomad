locals {
    fuzz_functions = [
        "FuzzMsgSwapWithinBatch_raw",
        "FuzzMsgSwapWithinBatch_structured",
        "FuzzMsgWithdrawWithinBatch_raw",
        "FuzzMsgWithdrawWithinBatch_structured",
    ]
}

job "fleece.liquidity" {
  region = "global"

  datacenters = ["dc1"]

  type = "service"

  group "x.liquidity.types" {
    count = 1

    dynamic "task" {
        labels = [task.value]
        for_each = local.fuzz_functions

        content {
            driver = "docker"
            config {
                image = "bryanchriswhite/go-fuzz:latest"
                entrypoint = ["/go-fuzz.sh"]
                args = [
                    "./x/liquidity/types", task.value,
                    "--", "-workdir", "./alloc/liquidity/fleece/workdirs/${task.value}",
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
}
