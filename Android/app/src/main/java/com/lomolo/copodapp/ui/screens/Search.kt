package com.lomolo.copodapp.ui.screens

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Search
import androidx.compose.material.icons.twotone.CheckCircle
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ListItem
import androidx.compose.material3.ListItemDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.SearchBar
import androidx.compose.material3.SearchBarDefaults
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.isTraversalGroup
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.traversalIndex
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.SearchLandViewModel
import com.lomolo.copodapp.state.viewmodels.SearchingLand
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.navigation.Navigation
import org.koin.androidx.compose.koinViewModel

object SearchScreenDestination : Navigation {
    override val title = null
    override val route = "search_land"
}

@Composable
fun SearchLandScreen(
    modifier: Modifier = Modifier,
    currentDestination: NavDestination,
    onNavigateTo: (String) -> Unit,
    onNavigateToFoundLand: (String) -> Unit,
    viewModel: SearchLandViewModel = koinViewModel<SearchLandViewModel>(),
) {
    Scaffold(topBar = {
        SearchLandTopBar(
            viewModel = viewModel,
            onNavigateToFoundLand = onNavigateToFoundLand,
        )
    }, bottomBar = {
        BottomNavBar(currentDestination = currentDestination, onNavigateTo = onNavigateTo)
    }) { innerPadding ->
        Surface(
            modifier = modifier
                .fillMaxSize()
                .padding(innerPadding)
        ) {}
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun SearchLandTopBar(
    modifier: Modifier = Modifier,
    viewModel: SearchLandViewModel,
    onNavigateToFoundLand: (String) -> Unit,
) {
    val searchQuery by viewModel.searchQuery.collectAsState()
    var expanded by rememberSaveable { mutableStateOf(false) }
    val searchResult by viewModel.searchResult.collectAsState()

    BackHandler {}
    Box(
        modifier
            .fillMaxSize()
            .semantics { isTraversalGroup = true }) {
        SearchBar(
            modifier = Modifier
                .align(Alignment.TopCenter)
                .semantics { traversalIndex = 0f },
            inputField = {
                SearchBarDefaults.InputField(query = searchQuery,
                    onSearch = { viewModel.searchLandTitle() },
                    expanded = expanded,
                    onExpandedChange = { expanded = it },
                    placeholder = { Text(stringResource(R.string.search_land)) },
                    leadingIcon = {
                        if (!expanded) {
                            Icon(
                                Icons.Default.Search, contentDescription = null
                            )
                        } else {
                            if (viewModel.searchingLand is SearchingLand.Loading) CircularProgressIndicator(
                                Modifier.size(16.dp)
                            ) else IconButton(onClick = {
                                expanded = false
                                viewModel.resetState()
                            }) {
                                Icon(
                                    Icons.AutoMirrored.Default.ArrowBack,
                                    contentDescription = stringResource(R.string.go_back)
                                )
                            }
                        }
                    },
                    onQueryChange = { viewModel.updateSearchQuery(it) })
            },
            onExpandedChange = { expanded = it },
            expanded = expanded,
        ) {
            // TODO: show list of found parcels
            Column(
                Modifier.padding(8.dp), verticalArrangement = Arrangement.spacedBy(4.dp)
            ) {
                if (searchResult) {
                    ListItem(headlineContent = {
                        Text(
                            searchQuery, style = MaterialTheme.typography.titleMedium
                        )
                    },
                        supportingContent = { Text(stringResource(R.string.found_this_title)) },
                        colors = ListItemDefaults.colors(containerColor = MaterialTheme.colorScheme.secondaryContainer),
                        leadingContent = {
                            Icon(
                                Icons.TwoTone.CheckCircle,
                                tint = MaterialTheme.colorScheme.primary,
                                modifier = Modifier.size(24.dp),
                                contentDescription = stringResource(R.string.check)
                            )
                        },
                        modifier = Modifier
                            .clickable {
                                onNavigateToFoundLand(FoundLandScreenDestination.route)
                            }
                            .fillMaxWidth()
                            .padding(horizontal = 4.dp, vertical = 4.dp))
                } else {
                    if (searchQuery.isNotEmpty()) {
                        Text(
                            stringResource(R.string.no_results_found),
                            style = MaterialTheme.typography.titleMedium,
                        )
                        Text(
                            stringResource(R.string.no_search_results)
                        )
                    }
                }
            }
        }
    }
}