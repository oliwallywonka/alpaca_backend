export function generateCode(name?: string) {
  const chars = 'abcdefghijklmnopqrstuvwxyz0123456789'
  const codeLength = 10
  const nameCodeLength = 4
  let code = ''

  if (!name) {
    for (let i = 0; i < codeLength; i++) {
      code += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    return code
  }

  const initCode = name.toUpperCase().slice(0, 2)
  for (let i = 0; i < nameCodeLength; i++) {
    code += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  return `${initCode}-${code}`
}

export function generateSlug(text: string) {
  return text.toLowerCase().replace(/\s+/g, '-')
}

export function slugify(str: string, delimiter = '_', preserved = ['.', '=', '-']) {
  if (str === '') {
    return ''
  }

  // special characters
  const specialCharsMap: { [key: string]: RegExp } = {
    a: /а|à|á|å|â/gi,
    b: /б/gi,
    c: /ц|ç/gi,
    d: /д/gi,
    e: /е|è|é|ê|ẽ|ë/gi,
    f: /ф/gi,
    g: /г/gi,
    h: /х/gi,
    i: /й|и|ì|í|î/gi,
    j: /ж/gi,
    k: /к/gi,
    l: /л/gi,
    m: /м/gi,
    n: /н|ñ/gi,
    o: /о|ò|ó|ô|ø/gi,
    p: /п/gi,
    q: /я/gi,
    r: /р/gi,
    s: /с/gi,
    t: /т/gi,
    u: /ю|ù|ú|ů|û/gi,
    v: /в/gi,
    w: /в/gi,
    x: /ь/gi,
    y: /ъ/gi,
    z: /з/gi,
    ae: /ä|æ/gi,
    oe: /ö/gi,
    ue: /ü/gi,
    Ae: /Ä/gi,
    Ue: /Ü/gi,
    Oe: /Ö/gi,
    ss: /ß/gi,
    and: /&/gi,
  }

  // replace special characters
  for (const k in specialCharsMap) {
    str = str.replace(specialCharsMap[k], k)
  }

  return str
    .replace(new RegExp('[' + preserved.join('') + ']', 'g'), ' ') // replace preserved characters with spaces
    .replace(/[^\w\ ]/gi, '') // replaces all non-alphanumeric with empty string
    .replace(/\s+/g, delimiter) // collapse whitespaces and replace with `delimiter`
}
