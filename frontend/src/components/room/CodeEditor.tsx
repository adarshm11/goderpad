import Editor from '@monaco-editor/react';
import { useState, useRef, useContext } from 'react';
import { DarkModeContext } from '../../App';
import { SandpackProvider, SandpackPreview } from '@codesandbox/sandpack-react';

function CodeEditor() {
  const { isDark } = useContext(DarkModeContext);
  const editorRef = useRef<any>(null);
  const [code, setCode] = useState('function App() {\n  return (\n    <div>\n      <h1>Hello, World!</h1>\n    </div>\n  );\n}\nexport default App;');

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

  const handleEditorDidMount = (editor: any, _: any) => {
    editorRef.current = editor;
  }

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setCode(value);
    }
  }

  return (
    <div className={`flex flex-row gap-4 ${isDark ? 'bg-slate-900' : 'bg-gray-100'} p-6 pt-20`}>
      <div className="border-2 border-white rounded-lg overflow-hidden w-1/2">
        <Editor 
          height="85vh" 
          defaultLanguage={"javascript"} 
          value={code} 
          theme={isDark ? 'slate-dark' : 'vs'}
          beforeMount={handleEditorWillMount}
          onMount={handleEditorDidMount}
          onChange={handleEditorChange}
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            padding: { top: 10 },
            scrollBeyondLastLine: false,
          }}
        />
      </div>
      <div className="w-1/2 border-2 border-white rounded-lg overflow-hidden h-[85vh]">
        <SandpackProvider
          template={"react"}
          files={{
            '/App.js': code,
          }}
          theme={isDark ? 'dark' : 'light'}
          options={{
            externalResources: [],
          }}
          style={{ height: '100%' }}
        >
          <SandpackPreview 
            style={{ height: '100%' }}
            showOpenInCodeSandbox={false}
          />
        </SandpackProvider>
      </div>
    </div>
  );
}

export default CodeEditor;
