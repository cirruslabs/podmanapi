load("github.com/cirrus-templates/golang@v0.1.0", "test_task")

def main(ctx):
    return [test_task()]
