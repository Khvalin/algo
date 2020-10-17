const traverse = function(node, callback) {
   callback(node);
   for (const child of node.childNodes) {
       traverse(child, callback);  // recursive call
   }
}