import React from 'react';

import type { Meta, StoryObj } from '@storybook/react';

import { Row } from './Row';

//👇 This default export determines where your story goes in the story list
const meta: Meta<typeof Row> = {
  component: Row,
};

export default meta;

type Story = StoryObj<typeof Row>;

export const Default: Story = {
  args: {
    variant: 'default',
  },
  render: args => <Row {...args}>Blah blah blah</Row>,
};
export const Selected: Story = {
  args: {
    variant: 'selected',
  },
  render: args => <Row {...args}>Blah blah blah</Row>,
};
export const Success: Story = {
  args: {
    variant: 'success',
  },
  render: args => <Row {...args}>Blah blah blah</Row>,
};
