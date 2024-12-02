import type { CodegenConfig } from '@graphql-codegen/cli'
 
const config: CodegenConfig = {
  schema: 'http://localhost:4545/graphql',
  documents: ['src/graphql/**/*.tsx'],
  ignoreNoDocuments: true,
  generates: {
    './src/graphql/': {
      preset: 'client',
      config: {
        documentMode: 'string'
      }
    },
  }
}
 
export default config
