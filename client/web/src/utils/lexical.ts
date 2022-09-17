export function extractMentionIds(obj: AnyObject): [number[], number[]] {
  const userIds = []
  const topicIds = []
  if (obj.children) {
    for (const child of obj.children) {
      if (child.text && child.type === 'mention' && child.mentionName) {
        if (child.text.startsWith('#'))
          topicIds.push(parseInt(child.mentionName))
        else if (child.text.startsWith('@'))
          userIds.push(parseInt(child.mentionName))
      }
      if (child.children) {
        const ids = extractMentionIds(child)
        topicIds.push(...ids[0])
        userIds.push(...ids[1])
      }
    }
  }
  return [topicIds, userIds]
}
