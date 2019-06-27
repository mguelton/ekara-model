package model

//*****************************************************************************
//
//     _         _           ____                           _           _
//    / \  _   _| |_ ___    / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| |
//   / _ \| | | | __/ _ \  | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` |
//  / ___ \ |_| | || (_) | | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| |
// /_/   \_\__,_|\__\___/   \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// ----------------------------------------------------
// Implementation(s) of TComponent
// ----------------------------------------------------

//TComponentOnDistributionHolder is the struct containing the Distribution in order to implement TComponent
type TComponentOnDistributionHolder struct {
	h Distribution
}

//CreateTComponentForDistribution returns an holder of Distribution implementing TComponent
func CreateTComponentForDistribution(o Distribution) TComponentOnDistributionHolder {
	return TComponentOnDistributionHolder{
		h: o,
	}
}

//ID returns the name of the component
func (r TComponentOnDistributionHolder) ID() string {
	return r.h.Id
}

//Repository returns the repository where the component is located
func (r TComponentOnDistributionHolder) Repository() TRepository {
	return CreateTRepositoryForRepository(r.h.Repository)
}

//HasTemplates returns true if the component has defined templates
func (r TComponentOnDistributionHolder) HasTemplates() bool {
	return len(r.h.Templates.Content) > 0
}

//Templates returns true if the component templates
func (r TComponentOnDistributionHolder) Templates() []string {
	return r.h.Templates.Content
}

//TComponentOnComponentHolder is the struct containing the Component in order to implement TComponent
type TComponentOnComponentHolder struct {
	h Component
}

//CreateTComponentForComponent returns an holder of Component implementing TComponent
func CreateTComponentForComponent(o Component) TComponentOnComponentHolder {
	return TComponentOnComponentHolder{
		h: o,
	}
}

//ID returns the name of the component
func (r TComponentOnComponentHolder) ID() string {
	return r.h.Id
}

//Repository returns the repository where the component is located
func (r TComponentOnComponentHolder) Repository() TRepository {
	return CreateTRepositoryForRepository(r.h.Repository)
}

//HasTemplates returns true if the component has defined templates
func (r TComponentOnComponentHolder) HasTemplates() bool {
	return len(r.h.Templates.Content) > 0
}

//Templates returns true if the component templates
func (r TComponentOnComponentHolder) Templates() []string {
	return r.h.Templates.Content
}