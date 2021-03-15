// Code generated from EGraphQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // EGraphQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEGraphQLListener is a complete listener for a parse tree produced by EGraphQLParser.
type BaseEGraphQLListener struct{}

var _ EGraphQLListener = &BaseEGraphQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEGraphQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEGraphQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEGraphQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEGraphQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseEGraphQLListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseEGraphQLListener) ExitDocument(ctx *DocumentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseEGraphQLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseEGraphQLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterExecutableDefinition is called when production executableDefinition is entered.
func (s *BaseEGraphQLListener) EnterExecutableDefinition(ctx *ExecutableDefinitionContext) {}

// ExitExecutableDefinition is called when production executableDefinition is exited.
func (s *BaseEGraphQLListener) ExitExecutableDefinition(ctx *ExecutableDefinitionContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *BaseEGraphQLListener) EnterOperationDefinition(ctx *OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *BaseEGraphQLListener) ExitOperationDefinition(ctx *OperationDefinitionContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *BaseEGraphQLListener) EnterOperationType(ctx *OperationTypeContext) {}

// ExitOperationType is called when production operationType is exited.
func (s *BaseEGraphQLListener) ExitOperationType(ctx *OperationTypeContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseEGraphQLListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseEGraphQLListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterSelection is called when production selection is entered.
func (s *BaseEGraphQLListener) EnterSelection(ctx *SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *BaseEGraphQLListener) ExitSelection(ctx *SelectionContext) {}

// EnterField is called when production field is entered.
func (s *BaseEGraphQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseEGraphQLListener) ExitField(ctx *FieldContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseEGraphQLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseEGraphQLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseEGraphQLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseEGraphQLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseEGraphQLListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseEGraphQLListener) ExitAlias(ctx *AliasContext) {}

// EnterFragmentSpread is called when production fragmentSpread is entered.
func (s *BaseEGraphQLListener) EnterFragmentSpread(ctx *FragmentSpreadContext) {}

// ExitFragmentSpread is called when production fragmentSpread is exited.
func (s *BaseEGraphQLListener) ExitFragmentSpread(ctx *FragmentSpreadContext) {}

// EnterFragmentDefinition is called when production fragmentDefinition is entered.
func (s *BaseEGraphQLListener) EnterFragmentDefinition(ctx *FragmentDefinitionContext) {}

// ExitFragmentDefinition is called when production fragmentDefinition is exited.
func (s *BaseEGraphQLListener) ExitFragmentDefinition(ctx *FragmentDefinitionContext) {}

// EnterFragmentName is called when production fragmentName is entered.
func (s *BaseEGraphQLListener) EnterFragmentName(ctx *FragmentNameContext) {}

// ExitFragmentName is called when production fragmentName is exited.
func (s *BaseEGraphQLListener) ExitFragmentName(ctx *FragmentNameContext) {}

// EnterTypeCondition is called when production typeCondition is entered.
func (s *BaseEGraphQLListener) EnterTypeCondition(ctx *TypeConditionContext) {}

// ExitTypeCondition is called when production typeCondition is exited.
func (s *BaseEGraphQLListener) ExitTypeCondition(ctx *TypeConditionContext) {}

// EnterInlineFragment is called when production inlineFragment is entered.
func (s *BaseEGraphQLListener) EnterInlineFragment(ctx *InlineFragmentContext) {}

// ExitInlineFragment is called when production inlineFragment is exited.
func (s *BaseEGraphQLListener) ExitInlineFragment(ctx *InlineFragmentContext) {}

// EnterValue is called when production value is entered.
func (s *BaseEGraphQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseEGraphQLListener) ExitValue(ctx *ValueContext) {}

// EnterIntValue is called when production intValue is entered.
func (s *BaseEGraphQLListener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production intValue is exited.
func (s *BaseEGraphQLListener) ExitIntValue(ctx *IntValueContext) {}

// EnterFloatValue is called when production floatValue is entered.
func (s *BaseEGraphQLListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production floatValue is exited.
func (s *BaseEGraphQLListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseEGraphQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseEGraphQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseEGraphQLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseEGraphQLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNullValue is called when production nullValue is entered.
func (s *BaseEGraphQLListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production nullValue is exited.
func (s *BaseEGraphQLListener) ExitNullValue(ctx *NullValueContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseEGraphQLListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseEGraphQLListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterListValue is called when production listValue is entered.
func (s *BaseEGraphQLListener) EnterListValue(ctx *ListValueContext) {}

// ExitListValue is called when production listValue is exited.
func (s *BaseEGraphQLListener) ExitListValue(ctx *ListValueContext) {}

// EnterObjectValue is called when production objectValue is entered.
func (s *BaseEGraphQLListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production objectValue is exited.
func (s *BaseEGraphQLListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterObjectField is called when production objectField is entered.
func (s *BaseEGraphQLListener) EnterObjectField(ctx *ObjectFieldContext) {}

// ExitObjectField is called when production objectField is exited.
func (s *BaseEGraphQLListener) ExitObjectField(ctx *ObjectFieldContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseEGraphQLListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseEGraphQLListener) ExitVariable(ctx *VariableContext) {}

// EnterVariableDefinitions is called when production variableDefinitions is entered.
func (s *BaseEGraphQLListener) EnterVariableDefinitions(ctx *VariableDefinitionsContext) {}

// ExitVariableDefinitions is called when production variableDefinitions is exited.
func (s *BaseEGraphQLListener) ExitVariableDefinitions(ctx *VariableDefinitionsContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseEGraphQLListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseEGraphQLListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseEGraphQLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseEGraphQLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseEGraphQLListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseEGraphQLListener) ExitType_(ctx *Type_Context) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseEGraphQLListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseEGraphQLListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseEGraphQLListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseEGraphQLListener) ExitListType(ctx *ListTypeContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseEGraphQLListener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseEGraphQLListener) ExitDirectives(ctx *DirectivesContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseEGraphQLListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseEGraphQLListener) ExitDirective(ctx *DirectiveContext) {}

// EnterTypeSystemDefinition is called when production typeSystemDefinition is entered.
func (s *BaseEGraphQLListener) EnterTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// ExitTypeSystemDefinition is called when production typeSystemDefinition is exited.
func (s *BaseEGraphQLListener) ExitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// EnterTypeSystemExtension is called when production typeSystemExtension is entered.
func (s *BaseEGraphQLListener) EnterTypeSystemExtension(ctx *TypeSystemExtensionContext) {}

// ExitTypeSystemExtension is called when production typeSystemExtension is exited.
func (s *BaseEGraphQLListener) ExitTypeSystemExtension(ctx *TypeSystemExtensionContext) {}

// EnterSchemaDefinition is called when production schemaDefinition is entered.
func (s *BaseEGraphQLListener) EnterSchemaDefinition(ctx *SchemaDefinitionContext) {}

// ExitSchemaDefinition is called when production schemaDefinition is exited.
func (s *BaseEGraphQLListener) ExitSchemaDefinition(ctx *SchemaDefinitionContext) {}

// EnterRootOperationTypeDefinition is called when production rootOperationTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterRootOperationTypeDefinition(ctx *RootOperationTypeDefinitionContext) {
}

// ExitRootOperationTypeDefinition is called when production rootOperationTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitRootOperationTypeDefinition(ctx *RootOperationTypeDefinitionContext) {
}

// EnterSchemaExtension is called when production schemaExtension is entered.
func (s *BaseEGraphQLListener) EnterSchemaExtension(ctx *SchemaExtensionContext) {}

// ExitSchemaExtension is called when production schemaExtension is exited.
func (s *BaseEGraphQLListener) ExitSchemaExtension(ctx *SchemaExtensionContext) {}

// EnterOperationTypeDefinition is called when production operationTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// ExitOperationTypeDefinition is called when production operationTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseEGraphQLListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseEGraphQLListener) ExitDescription(ctx *DescriptionContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseEGraphQLListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseEGraphQLListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterTempletableTypeDefinition is called when production templetableTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterTempletableTypeDefinition(ctx *TempletableTypeDefinitionContext) {
}

// ExitTempletableTypeDefinition is called when production templetableTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitTempletableTypeDefinition(ctx *TempletableTypeDefinitionContext) {}

// EnterTypeExtension is called when production typeExtension is entered.
func (s *BaseEGraphQLListener) EnterTypeExtension(ctx *TypeExtensionContext) {}

// ExitTypeExtension is called when production typeExtension is exited.
func (s *BaseEGraphQLListener) ExitTypeExtension(ctx *TypeExtensionContext) {}

// EnterScalarTypeDefinition is called when production scalarTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// ExitScalarTypeDefinition is called when production scalarTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// EnterScalarTypeExtension is called when production scalarTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterScalarTypeExtension(ctx *ScalarTypeExtensionContext) {}

// ExitScalarTypeExtension is called when production scalarTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitScalarTypeExtension(ctx *ScalarTypeExtensionContext) {}

// EnterObjectTypeDefinition is called when production objectTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// ExitObjectTypeDefinition is called when production objectTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// EnterImplementsInterfaces is called when production implementsInterfaces is entered.
func (s *BaseEGraphQLListener) EnterImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// ExitImplementsInterfaces is called when production implementsInterfaces is exited.
func (s *BaseEGraphQLListener) ExitImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// EnterFieldsDefinition is called when production fieldsDefinition is entered.
func (s *BaseEGraphQLListener) EnterFieldsDefinition(ctx *FieldsDefinitionContext) {}

// ExitFieldsDefinition is called when production fieldsDefinition is exited.
func (s *BaseEGraphQLListener) ExitFieldsDefinition(ctx *FieldsDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseEGraphQLListener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseEGraphQLListener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterArgumentsDefinition is called when production argumentsDefinition is entered.
func (s *BaseEGraphQLListener) EnterArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// ExitArgumentsDefinition is called when production argumentsDefinition is exited.
func (s *BaseEGraphQLListener) ExitArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// EnterInputValueDefinition is called when production inputValueDefinition is entered.
func (s *BaseEGraphQLListener) EnterInputValueDefinition(ctx *InputValueDefinitionContext) {}

// ExitInputValueDefinition is called when production inputValueDefinition is exited.
func (s *BaseEGraphQLListener) ExitInputValueDefinition(ctx *InputValueDefinitionContext) {}

// EnterObjectTypeExtension is called when production objectTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterObjectTypeExtension(ctx *ObjectTypeExtensionContext) {}

// ExitObjectTypeExtension is called when production objectTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitObjectTypeExtension(ctx *ObjectTypeExtensionContext) {}

// EnterInterfaceTypeDefinition is called when production interfaceTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// ExitInterfaceTypeDefinition is called when production interfaceTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// EnterInterfaceTypeExtension is called when production interfaceTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) {}

// ExitInterfaceTypeExtension is called when production interfaceTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) {}

// EnterUnionTypeDefinition is called when production unionTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// ExitUnionTypeDefinition is called when production unionTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// EnterUnionMemberTypes is called when production unionMemberTypes is entered.
func (s *BaseEGraphQLListener) EnterUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// ExitUnionMemberTypes is called when production unionMemberTypes is exited.
func (s *BaseEGraphQLListener) ExitUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// EnterUnionTypeExtension is called when production unionTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterUnionTypeExtension(ctx *UnionTypeExtensionContext) {}

// ExitUnionTypeExtension is called when production unionTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitUnionTypeExtension(ctx *UnionTypeExtensionContext) {}

// EnterEnumTypeDefinition is called when production enumTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// ExitEnumTypeDefinition is called when production enumTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// EnterEnumValuesDefinition is called when production enumValuesDefinition is entered.
func (s *BaseEGraphQLListener) EnterEnumValuesDefinition(ctx *EnumValuesDefinitionContext) {}

// ExitEnumValuesDefinition is called when production enumValuesDefinition is exited.
func (s *BaseEGraphQLListener) ExitEnumValuesDefinition(ctx *EnumValuesDefinitionContext) {}

// EnterEnumValueDefinition is called when production enumValueDefinition is entered.
func (s *BaseEGraphQLListener) EnterEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// ExitEnumValueDefinition is called when production enumValueDefinition is exited.
func (s *BaseEGraphQLListener) ExitEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// EnterEnumTypeExtension is called when production enumTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterEnumTypeExtension(ctx *EnumTypeExtensionContext) {}

// ExitEnumTypeExtension is called when production enumTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitEnumTypeExtension(ctx *EnumTypeExtensionContext) {}

// EnterInputObjectTypeDefinition is called when production inputObjectTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {
}

// ExitInputObjectTypeDefinition is called when production inputObjectTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {}

// EnterInputFieldsDefinition is called when production inputFieldsDefinition is entered.
func (s *BaseEGraphQLListener) EnterInputFieldsDefinition(ctx *InputFieldsDefinitionContext) {}

// ExitInputFieldsDefinition is called when production inputFieldsDefinition is exited.
func (s *BaseEGraphQLListener) ExitInputFieldsDefinition(ctx *InputFieldsDefinitionContext) {}

// EnterInputObjectTypeExtension is called when production inputObjectTypeExtension is entered.
func (s *BaseEGraphQLListener) EnterInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) {}

// ExitInputObjectTypeExtension is called when production inputObjectTypeExtension is exited.
func (s *BaseEGraphQLListener) ExitInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) {}

// EnterDirectiveDefinition is called when production directiveDefinition is entered.
func (s *BaseEGraphQLListener) EnterDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// ExitDirectiveDefinition is called when production directiveDefinition is exited.
func (s *BaseEGraphQLListener) ExitDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// EnterDirectiveLocations is called when production directiveLocations is entered.
func (s *BaseEGraphQLListener) EnterDirectiveLocations(ctx *DirectiveLocationsContext) {}

// ExitDirectiveLocations is called when production directiveLocations is exited.
func (s *BaseEGraphQLListener) ExitDirectiveLocations(ctx *DirectiveLocationsContext) {}

// EnterDirectiveLocation is called when production directiveLocation is entered.
func (s *BaseEGraphQLListener) EnterDirectiveLocation(ctx *DirectiveLocationContext) {}

// ExitDirectiveLocation is called when production directiveLocation is exited.
func (s *BaseEGraphQLListener) ExitDirectiveLocation(ctx *DirectiveLocationContext) {}

// EnterExecutableDirectiveLocation is called when production executableDirectiveLocation is entered.
func (s *BaseEGraphQLListener) EnterExecutableDirectiveLocation(ctx *ExecutableDirectiveLocationContext) {
}

// ExitExecutableDirectiveLocation is called when production executableDirectiveLocation is exited.
func (s *BaseEGraphQLListener) ExitExecutableDirectiveLocation(ctx *ExecutableDirectiveLocationContext) {
}

// EnterTypeSystemDirectiveLocation is called when production typeSystemDirectiveLocation is entered.
func (s *BaseEGraphQLListener) EnterTypeSystemDirectiveLocation(ctx *TypeSystemDirectiveLocationContext) {
}

// ExitTypeSystemDirectiveLocation is called when production typeSystemDirectiveLocation is exited.
func (s *BaseEGraphQLListener) ExitTypeSystemDirectiveLocation(ctx *TypeSystemDirectiveLocationContext) {
}

// EnterName is called when production name is entered.
func (s *BaseEGraphQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseEGraphQLListener) ExitName(ctx *NameContext) {}

// EnterTemplateParametersDefinition is called when production templateParametersDefinition is entered.
func (s *BaseEGraphQLListener) EnterTemplateParametersDefinition(ctx *TemplateParametersDefinitionContext) {
}

// ExitTemplateParametersDefinition is called when production templateParametersDefinition is exited.
func (s *BaseEGraphQLListener) ExitTemplateParametersDefinition(ctx *TemplateParametersDefinitionContext) {
}

// EnterTemplateTypeDefinition is called when production templateTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterTemplateTypeDefinition(ctx *TemplateTypeDefinitionContext) {}

// ExitTemplateTypeDefinition is called when production templateTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitTemplateTypeDefinition(ctx *TemplateTypeDefinitionContext) {}

// EnterTemplateImplementedTypeDefinition is called when production templateImplementedTypeDefinition is entered.
func (s *BaseEGraphQLListener) EnterTemplateImplementedTypeDefinition(ctx *TemplateImplementedTypeDefinitionContext) {
}

// ExitTemplateImplementedTypeDefinition is called when production templateImplementedTypeDefinition is exited.
func (s *BaseEGraphQLListener) ExitTemplateImplementedTypeDefinition(ctx *TemplateImplementedTypeDefinitionContext) {
}
