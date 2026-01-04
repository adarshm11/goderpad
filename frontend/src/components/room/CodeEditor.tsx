import Editor from '@monaco-editor/react';
import { useContext } from 'react';
import { DarkModeContext } from '../../App';
import { languageToCode } from '../../util/languageToCode';

function CodeEditor({ language }: { language: string }) {
  const { isDark } = useContext(DarkModeContext);

  const handleEditorWillMount = (monaco: any) => {
    monaco.editor.defineTheme('slate-dark', {
      base: 'vs-dark',
      inherit: true,
      rules: [],
      colors: {
        'editor.background': '#0f172a', // slate-900
      },
    });
  };

  return (
    <div className={`${isDark ? 'bg-slate-900' : 'bg-gray-100'} p-6 pt-20`}>
      <div className="border-2 border-white rounded-lg overflow-hidden">
        <Editor 
          height="85vh" 
          language={language === 'c++' ? 'cpp' : language} 
          value={languageToCode(language)} 
          theme={isDark ? 'slate-dark' : 'vs'}
          beforeMount={handleEditorWillMount}
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            padding: { top: 10 },
            scrollBeyondLastLine: false,
          }}
        />
      </div>
    </div>
  );
}

export default CodeEditor;
