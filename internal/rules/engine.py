import yaml

def evaluate(state):
    with open("rules.yaml") as f:
        rules = yaml.safe_load(f)["rules"]
    return [r["action"] for r in rules if eval(r["condition"], {}, state)]
