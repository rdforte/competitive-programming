const maxProbability = function (n, edges, succProb, start_node, end_node) {
  const adjList = Array(n)
    .fill()
    .map(() => []);
  const probabilities = Array(n).fill(Number.NEGATIVE_INFINITY);
  probabilities[start_node] = 1;

  for (let i = 0; i < edges.length; i++) {
    const vertice1 = edges[i][0];
    const vertice2 = edges[i][1];

    adjList[vertice1].push({ ver: vertice2, prob: succProb[i] });
    adjList[vertice2].push({ ver: vertice1, prob: succProb[i] });
  }

  const mHeap = new MaxHeap();

  mHeap.add({ ver: start_node, prob: 1 });

  while (!mHeap.isEmpty()) {
    const node = mHeap.top();
    mHeap.pop();

    if (node.ver === end_node) return node.prob;

    for (let n of adjList[node.ver]) {
      const prob = node.prob * n.prob;
      const ver = n.ver;
      if (prob > probabilities[ver]) {
        probabilities[ver] = prob;
        mHeap.add({ prob, ver });
      }
    }

    adjList[node.ver] = [];
  }

  return 0;
};

const n = 5;
const edges = [
  [1, 4],
  [2, 4],
  [0, 4],
  [0, 3],
  [0, 2],
  [2, 3],
];
const succProb = [0.37, 0.17, 0.93, 0.23, 0.39, 0.04];
const start = 3;
const end = 4;

class MaxHeap {
  #heap = [];

  constructor() {}

  add(item) {
    this.#heap.push(item);
    this.#heap.sort((a, b) => a.prob - b.prob);
  }

  isEmpty() {
    return this.#heap.length === 0;
  }

  top() {
    return this.#heap[this.#heap.length - 1];
  }

  pop() {
    this.#heap.pop();
  }
}

console.log(maxProbability(n, edges, succProb, start, end));
