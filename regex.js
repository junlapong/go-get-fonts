//                                            Group[0]
//                                            vvvvvvvv
fontFaceRulePattern = /(?:@font-face\s*\{\s*)([^\{\}]+)/gi,
//
//                             Group[1]                               Group[2]
//                             vvvvvv                                 vvvvvv
fontUrlPattern = /url\([\'\"]?([^\)]+?)[\'\"]?\)(?:\s*format\([\'\"]?([^\)]+?)[\'\"]?\))?/gi,
//
//                                             Group[1]
//                   Group[0]                  vvvvvvv
//                   vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
fontFamilyPattern = /font-family\s*\:\s*[\'"]?([^\;]+?)[\'"]?\;/i,
fontWeightPattern = /font-weight\s*\:\s*[\'"]?([^\;]+?)[\'"]?\;/i,
fontStylePattern =   /font-style\s*\:\s*[\'"]?([^\;]+?)[\'"]?\;/i



const REX_URL = /url\s*\(\s*(?:\'|")?\s*([^]*?)\s*(?:\'|")?\s*\)\s*?/gi


const REGULAR_EXPRESSIONS = {
	face:   /\s*(?:\/\*\s*(.*?)\s*\*\/)?[^@]*?@font-face\s*{(?:[^}]*?)}\s*/gi,
	family: /font-family\s*:\s*(?:\'|")?([^;]*?)(?:\'|")?\s*;/i,
	weight: /font-weight\s*:\s*([^;]*?)\s*;/i,
	url:    /url\s*\(\s*(?:\'|")?\s*([^]*?)\s*(?:\'|")?\s*\)\s*?/gi
}