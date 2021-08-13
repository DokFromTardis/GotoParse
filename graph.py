import networkx as nx
import matplotlib.pyplot as plt
import json
path = 'members.json'
graph = nx.Graph()
with open(path, 'r', encoding='utf-8')as file:
    data = json.loads(file.read())
    for i in range(30):
      user = data[0]["response"][0]["user"]
      graph.add_node(user)
      for j in data[0]["response"][0]["friends"]:
        graph.add_edge(user, j["id"])
pos = nx.spring_layout(graph,
                       k=None,
                       pos=None,
                       fixed=None,
                       iterations=50,
                       threshold=0.001,
                       weight='weight',
                       scale=None,
                       center=None,
                       dim=2,
                       seed=None)
print(pos)
nx.draw_networkx(graph, pos)
plt.show()
plt.savefig("path.png")
