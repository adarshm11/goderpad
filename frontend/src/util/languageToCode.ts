
export function languageToCode(language: string): string {
  switch (language.toLowerCase()) {
    case 'python':
      return 'def solution():';
    case 'javascript':
      return 'const solution = () => {\n\n};';
    case 'java':
      return 'public class Solution {\n    public static void main(String[] args) {\n\n    }\n}';
    case 'c++':
      return 'void solution() {\n\n}';
    default:
      return '';
  }
}