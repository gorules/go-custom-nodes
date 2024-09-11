import { createJdmNode } from '@gorules/jdm-editor';
import { PlusOutlined } from '@ant-design/icons';

export const addNode = createJdmNode({
  kind: 'mathAdd' as string,
  group: 'math',
  icon: <PlusOutlined />,
  displayName: 'Add',
  shortDescription: 'Adds 2 numbers',
  inputs: [
    { label: 'a', name: 'a', control: 'text' },
    { label: 'b', name: 'b', control: 'text' },
    { label: 'Key', name: 'key', control: 'text' },
  ]
});