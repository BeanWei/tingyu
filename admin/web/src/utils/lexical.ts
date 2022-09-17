export function extractText(obj) {
  const output = []
  if (obj.children) {
    for (const child of obj.children) {
      if (child.text)
        output.push(child.text)
      if (child.children)
        output.push(extractText(child))
    }
  }
  return output.join('')
}
